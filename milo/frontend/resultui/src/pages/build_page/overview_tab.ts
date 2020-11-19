// Copyright 2020 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { MobxLitElement } from '@adobe/lit-mobx';
import '@material/mwc-button';
import '@material/mwc-dialog';
import '@material/mwc-textarea';
import { TextArea } from '@material/mwc-textarea';
import { Router } from '@vaadin/router';
import { css, customElement, html } from 'lit-element';
import { observable } from 'mobx';

import '../../components/build_step_entry';
import '../../components/code_mirror_editor';
import '../../components/link';
import { AppState, consumeAppState } from '../../context/app_state/app_state';
import { BuildState, consumeBuildState } from '../../context/build_state/build_state';
import { getBotLink, getBuildbucketLink, getURLForBuild, getURLForGerritChange, getURLForGitilesCommit, getURLForSwarmingTask } from '../../libs/build_utils';
import { BUILD_STATUS_CLASS_MAP, BUILD_STATUS_DISPLAY_MAP } from '../../libs/constants';
import { displayDuration, LONG_TIME_FORMAT } from '../../libs/time_utils';
import { renderMarkdown } from '../../libs/utils';
import { StepExt } from '../../models/step_ext';
import { router } from '../../routes';
import { BuildStatus } from '../../services/buildbucket';

export class OverviewTabElement extends MobxLitElement {
  @observable.ref appState!: AppState;
  @observable.ref buildState!: BuildState;

  @observable.ref private showRetryDialog = false;
  @observable.ref private showCancelDialog = false;

  connectedCallback() {
    super.connectedCallback();
    this.appState.selectedTabId = 'overview';
  }

  private renderStatusTime() {
    const build = this.buildState.build!;

    return html`
      <div id="status">
        Build
        <i class="status ${BUILD_STATUS_CLASS_MAP[build.status]}">
          ${BUILD_STATUS_DISPLAY_MAP[build.status] || 'unknown status'}
        </i>
        ${(() => { switch (build.status) {
        case BuildStatus.Scheduled:
          return `since ${build.createTime.toFormat(LONG_TIME_FORMAT)}`;
        case BuildStatus.Started:
          return `since ${build.startTime!.toFormat(LONG_TIME_FORMAT)}`;
        case BuildStatus.Canceled:
          return `after ${displayDuration(build.endTime!.diff(build.createTime))} by ${build.canceledBy}`;
        case BuildStatus.Failure:
        case BuildStatus.InfraFailure:
        case BuildStatus.Success:
          return `after ${displayDuration(build.endTime!.diff(build.startTime || build.createTime))}`;
        default:
          return '';
        }})()}
        ${build.endTime ?
          html`<mwc-button dense unelevated @click=${() => this.showRetryDialog = true}>Retry</mwc-button>` :
          html`<mwc-button dense unelevated @click=${() => this.showCancelDialog = true}>Cancel</mwc-button>`}
      </div>
    `;
  }

  private renderCanaryWarning() {
    if (!this.buildState.build?.isCanary) {
      return html``;
    }
    if ([BuildStatus.Failure, BuildStatus.InfraFailure].indexOf(this.buildState.build!.status) === -1) {
      return html``;
    }
    return html`
      <div id="canary-warning">
        WARNING: This build ran on a canary version of LUCI. If you suspect it
        failed due to infra, retry the build. Next time it may use the
        non-canary version.
      </div>
    `;
  }

  private async cancelBuild(reason: string) {
    await this.appState.buildsService!.cancelBuild({
      id: this.buildState.build!.id,
      summaryMarkdown: reason,
    });
    this.buildState.refresh();
  }

  private async retryBuild() {
    const build = await this.appState.buildsService!.scheduleBuild({
      templateBuildId: this.buildState.build!.id,
    });
    Router.go(getURLForBuild(build));
  }

  private renderSummary() {
    const build = this.buildState.build!;
    if (!build.summaryMarkdown) {
      return html``;
    }

    return html`
      <div id="summary-html">
        ${renderMarkdown(build.summaryMarkdown)}
      </div>
    `;
  }

  private renderInput() {
    const input = this.buildState.build?.input;
    if (!input) {
      return html``;
    }
    return html`
      <div>
        <h3>Input</h3>
        <table>
          ${input.gitilesCommit ? html`
            <tr>
              <td>Revision:</td>
              <td>
                <a href=${getURLForGitilesCommit(input.gitilesCommit)} target="_blank">${input.gitilesCommit.id}</a>
                ${input.gitilesCommit.position ? `CP #${input.gitilesCommit.position}` : ''}
              </td>
            </tr>
          ` : ''}
          ${(input.gerritChanges || []).map((gc) => html`
            <tr>
              <td>Patch:</td>
              <td>
                <a href=${getURLForGerritChange(gc)}>
                  ${gc.change} (ps #${gc.patchset})
                </a>
              </td>
            </tr>
          `)}
        </table>
      </div>
    `;
  }

  private renderInfra() {
    const build = this.buildState.build!;
    const botLink = build.infra?.swarming ? getBotLink(build.infra.swarming) : null;
    return html`
      <div>
        <h3>Infra</h3>
        <table>
          <tr><td>Buildbucket ID:</td><td><milo-link .link=${getBuildbucketLink(CONFIGS.BUILDBUCKET.HOST, build.id)} target="_blank"></td></tr>
          ${build.infra?.swarming ? html`
          <tr>
            <td>Swarming Task:</td>
            <td>${build.infra.swarming.taskId ? html`<a href=${getURLForSwarmingTask(build.infra.swarming)}>${build.infra.swarming.taskId}</a>`: 'N/A'}</td>
          </tr>
          <tr>
            <td>Bot:</td>
            <td>${botLink ? html`<milo-link .link=${botLink} target="_blank"></milo-link>` : 'N/A'}</td>
          </tr>
          ` : ''}
          <tr><td>Recipe:</td><td><milo-link .link=${build.recipeLink} target="_blank"></milo-link></td></tr>
        </table>
      </div>
    `;
  }

  private renderSteps() {
    const build = this.buildState.build!;
    const nonSucceededSteps = (build.rootSteps || [])
      .map((step, i) => [step, i + 1] as [StepExt, number])
      .filter(([step, _stepNum]) => !step.succeededRecursively);
    const scheduledSteps = nonSucceededSteps
      .filter(([step, _stepNum]) => step.status === BuildStatus.Scheduled);
    const runningSteps = nonSucceededSteps
      .filter(([step, _stepNum]) => step.status === BuildStatus.Started);
    const canceledSteps = nonSucceededSteps
      .filter(([step, _stepNum]) => step.status === BuildStatus.Canceled);
    const failedSteps = nonSucceededSteps
      .filter(([step, _stepNum]) => step.failed);

    return html`
      <div>
        <h3>Steps & Logs
          (<a href=${router.urlForName('build-steps', {
            ...this.buildState.builder,
            build_num_or_id: this.buildState.buildNumOrId!,
          })}>View All</a>)
        </h3>
        <div class="step-summary-line">
          ${this.renderStepSummary(failedSteps.length, canceledSteps.length, scheduledSteps.length, runningSteps.length)}
        </div>
        ${nonSucceededSteps.map(([step, stepNum]) => html`
        <milo-build-step-entry
          .expanded=${true}
          .number=${stepNum}
          .step=${step}
          .showDebugLogs=${false}
        ></milo-build-step-entry>
        `) || ''}
      </div>
    `;
  }

  private renderStepSummary(failedSteps: number, canceledSteps: number, scheduledSteps: number, runningSteps: number) {
    if (failedSteps === 0 && scheduledSteps === 0 && runningSteps === 0) {
        return 'All steps succeeded.';
    }
    const messageParts: string[] = [];
    if (failedSteps > 0) {
        messageParts.push(`${failedSteps} step${failedSteps === 1 ? '' : 's'} failed`);
    }
    if (canceledSteps > 0) {
        messageParts.push(`${canceledSteps} step${canceledSteps === 1 ? '' : 's'} canceled`);
    }
    if (scheduledSteps > 0) {
        messageParts.push(`${scheduledSteps} step${scheduledSteps === 1 ? '' : 's'} scheduled`);
    }
    if (runningSteps > 0) {
        messageParts.push(`${runningSteps} step${runningSteps === 1 ? '' : 's'} still running`);
    }
    return messageParts.join(', ') + ':';
  }

  private renderTiming() {
    const build = this.buildState.build!;

    return html`
      <div>
        <h3>Timing</h3>
        <table>
          <tr><td>Created:</td><td>${build.createTime.toFormat(LONG_TIME_FORMAT)} (${displayDuration(build.timeSinceCreated)} ago)</td></tr>
          <tr><td>Started:</td><td>${build.startTime && (build.startTime.toFormat(LONG_TIME_FORMAT) + ` (${displayDuration(build.timeSinceStarted!)} ago)`) || 'N/A'}</td></tr>
          <tr><td>Ended:</td><td>${build.endTime && (build.endTime.toFormat(LONG_TIME_FORMAT) + ` (${displayDuration(build.timeSinceEnded!)} ago)`) || 'N/A'}</td></tr>
          <tr><td>Pending:</td><td>${build.pendingDuration && displayDuration(build.pendingDuration) || 'N/A'}</td></tr>
          <tr><td>Execution:</td><td>${build.executionDuration && displayDuration(build.executionDuration) || 'N/A'}</td></tr>
        </table>
      </div>
    `;
  }

  private renderTags() {
    const tags = this.buildState.build?.tags;
    if (!tags) {
      return html``;
    }
    return html`
      <div>
        <h3>Tags</h3>
        <div id="tags">
          ${tags.map((tag) => html`
          <div class="key">${tag.key}:</div>
          <div class="value">${tag.value}</div>
          `)}
        </div>
      </div>
    `;
  }

  private renderProperties(header: string, properties: {[key: string]: unknown}) {
    const editorOptions = {
      mode: {name: 'javascript', json: true},
      readOnly: true,
      scrollbarStyle: 'null',
      matchBrackets: true,
      lineWrapping: true,
    };

    return html`
      <div>
        <h3>${header}</h3>
        <milo-code-mirror-editor
          .value=${JSON.stringify(properties, undefined, 2)}
          .options=${{...editorOptions}}
        ></milo-code-mirror-editor>
      </div>
    `;
  }

  protected render() {
    const build = this.buildState.build;
    if (!build) {
      return html``;
    }

    return html`
      <mwc-dialog
        heading="Retry Build"
        ?open=${this.showRetryDialog}
        @closed=${async (event: CustomEvent<{action: string}>) => {
          if (event.detail.action === 'retry') {
            await this.retryBuild();
          }
          this.showRetryDialog = false;
        }}
      >
        <p>Note: this doesn't trigger anything else (e.g. CQ).</p>
        <mwc-button slot="primaryAction" dialogAction="retry" dense unelevated>Retry</mwc-button>
        <mwc-button slot="secondaryAction" dialogAction="dismiss">Dismiss</mwc-button>
      </mwc-dialog>
      <mwc-dialog
        heading="Cancel Build"
        ?open=${this.showCancelDialog}
        @closed=${async (event: CustomEvent<{action: string}>) => {
          if (event.detail.action === 'cancel') {
            const reason = (this.shadowRoot!.getElementById('cancel-reason') as TextArea).value;
            await this.cancelBuild(reason);
          }
          this.showCancelDialog = false;
        }}
      >
        <mwc-textarea id="cancel-reason" label="Reason" required></mwc-textarea>
        <mwc-button slot="primaryAction" dialogAction="cancel" dense unelevated>Cancel</mwc-button>
        <mwc-button slot="secondaryAction" dialogAction="dismiss">Dismiss</mwc-button>
      </mwc-dialog>
      <div id="main">
        <div class="first-column">
          ${this.renderStatusTime()}
          ${this.renderCanaryWarning()}
          ${this.renderSummary()}
          ${this.renderInput()}
          ${this.renderInfra()}
          ${this.renderTiming()}
          <!-- TODO(crbug/1116824): render failed tests -->
          ${this.renderSteps()}
        </div>
        <div class="second-column">
          ${this.renderTags()}
          ${this.renderProperties('Input Properties', build.input.properties)}
          ${this.renderProperties('Output Properties', build.output.properties)}
        </div>
      </div>
    `;
  }

  static styles = css`
    #main {
      margin: 5px 24px;
    }
    @media screen and (min-width: 1500px) {
      #main {
        display: grid;
        grid-template-columns: auto 1fr;
        grid-gap: 20px;
      }
      .first-column {
        max-width: 900px;
      }
      .second-column {
        overflow: hidden;
      }
    }

    #status {
      font-weight: 500;
    }
    .status.scheduled {
      color: var(--scheduled-color);
    }
    .status.started {
      color: var(--started-color);
    }
    .status.success {
      color: var(--success-color);
    }
    .status.failure {
      color: var(--failure-color);
    }
    .status.infra-failure {
      color: var(--critical-failure-color);
    }
    .status.canceled {
      color: var(--canceled-color);
    }

    :host > mwc-dialog {
      margin: 0 0;
    }
    #cancel-reason {
      width: 500px;
      height: 200px;
    }
    mwc-button {
      transform: scale(0.8);
      vertical-align: middle;
    }

    #canary-warning {
      background-color: var(--warning-color);
      font-weight: 500;
    }

    #summary-html {
      background-color: var(--block-background-color);
      padding: 5px;
      clear: both;
      overflow-wrap: break-word;
    }

    #summary-html pre {
      white-space: pre-wrap;
    }

    #tags {
      display: grid;
      grid-template-columns: auto 1fr;
    }
    #tags div {
      clear: both;
      overflow-wrap: anywhere;
    }
    #tags .value {
      margin-left: 5px;
    }
    #tags>div {
      margin-top: 1px;
      margin-bottom: 1px;
    }

    .step-summary-line {
      margin-bottom: 10px;
    }

    milo-code-mirror-editor {
      min-width: 600px;
      max-width: 1000px;
    }
  `;
}

customElement('milo-overview-tab')(
  consumeBuildState(
    consumeAppState(OverviewTabElement),
  ),
);
