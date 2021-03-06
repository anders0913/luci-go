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
import '@material/mwc-icon';
import { css, customElement, html } from 'lit-element';
import { styleMap } from 'lit-html/directives/style-map';
import { computed, observable } from 'mobx';
import { fromPromise, FULFILLED, IPromiseBasedObservable } from 'mobx-utils';

import { AppState, consumeAppState } from '../../context/app_state/app_state';
import '../../context/artifact/artifact_provider';
import { TEST_STATUS_DISPLAY_MAP } from '../../libs/constants';
import { sanitizeHTML } from '../../libs/sanitize_html';
import { ListArtifactsResponse, TestResult } from '../../services/resultdb';
import '../expandable_entry';
import './image_diff_artifact';
import './text_artifact';
import './text_diff_artifact';

/**
 * Renders an expandable entry of the given test result.
 */
export class ResultEntryElement extends MobxLitElement {
  @observable.ref id = '';
  @observable.ref testResult!: TestResult;
  @observable.ref appState!: AppState;

  @observable.ref private _expanded = false;
  @computed get expanded() { return this._expanded; }
  set expanded(newVal: boolean) {
    this._expanded = newVal;
    // Always render the content once it was expanded so the descendants' states
    // don't get reset after the node is collapsed.
    this.shouldRenderContent = this.shouldRenderContent || newVal;
  }

  @observable.ref private shouldRenderContent = false;

  @observable.ref private tagExpanded = false;

  @computed
  private get parentInvId() {
    return /^invocations\/(.+?)\/.+$/.exec(this.testResult.name)![1];
  }

  @computed
  private get parentInvType() {
    if (this.parentInvId.startsWith('build-')) {
      return 'Buildbucket build';
    }
    if (this.parentInvId.startsWith('task-')) {
      return 'Swarming task';
    }
    return 'invocation';
  }

  @computed
  private get resultArtifactsRes(): IPromiseBasedObservable<ListArtifactsResponse> {
    if (!this.appState.resultDb) {
      // Returns a promise that never resolves when resultDb isn't ready.
      return fromPromise(Promise.race([]));
    }
    // TODO(weiweilin): handle pagination.
    return fromPromise(this.appState.resultDb.listArtifacts({parent: this.testResult.name}));
  }

  @computed private get resultArtifacts() {
    return this.resultArtifactsRes.state === FULFILLED
      ? this.resultArtifactsRes.value.artifacts || []
      : [];
  }

  @computed private get invArtifacts() {
    if (!this.appState.resultDb) {
      return [];
    }
    return this.appState.resultDb.getCachedArtifactsOfInv('invocations/' + this.parentInvId) || [];
  }

  @computed private get artifactsMapping() {
    return new Map(this.resultArtifacts.map(obj => [obj.artifactId, obj]));
  }

  @computed private get textDiffArtifact() {
    return this.resultArtifacts.find((a) => a.artifactId === 'text_diff');
  }
  @computed private get imageDiffArtifactGroup() {
    return {
      'expected': this.resultArtifacts.find((a) => a.artifactId === 'expected_image'),
      'actual': this.resultArtifacts.find((a) => a.artifactId === 'actual_image'),
      'diff': this.resultArtifacts.find((a) => a.artifactId === 'image_diff'),
    };
  }

  private renderSummaryHtml() {
    if (!this.testResult.summaryHtml) {
      return html``;
    }

    return html`
    <div id="summary-html">
      <milo-artifact-provider .artifacts=${this.artifactsMapping}>
        ${sanitizeHTML(this.testResult.summaryHtml)}
      </milo-artifact-provider>
    </div>
    `;
  }

  private renderTags() {
    if ((this.testResult.tags || []).length === 0) {
      return html``;
    }

    return html`
      <milo-expandable-entry .hideContentRuler=${true}
        .onToggle=${(expanded: boolean) => {
          this.tagExpanded = expanded;
        }}
      >
        <span slot="header" class="one-line-content">
          Tags:
          <span class="greyed-out" style=${styleMap({display: this.tagExpanded ? 'none': ''})}>
            ${this.testResult.tags?.map((tag) => html`
            <span class="kv-key">${tag.key}</span>
            <span class="kv-value">${tag.value}</span>
            `)}
          </span>
        </span>
        <table id="tag-table" slot="content" border="0">
          ${this.testResult.tags?.map((tag) => html`
          <tr>
            <td>${tag.key}:</td>
            <td>${tag.value}</td>
          </tr>
          `)}
        </table>
      </milo-expandable-entry>
    `;
  }

  private renderInvocationLevelArtifacts() {
    if (this.invArtifacts.length === 0) {
      return html``;
    }

    return html`
      <div id="inv-artifacts-header">
        From the parent ${this.parentInvType}:
      </div>
      <ul>
        ${this.invArtifacts.map((artifact) => html`
        <!-- TODO(weiweilin): refresh when the fetchUrl expires -->
        <li>
          <a href=${artifact.fetchUrl} target="_blank">${artifact.artifactId}</a>
        </li>
        `)}
      </ul>
    `;
  }

  private renderArtifacts() {
    const artifactCount = this.resultArtifacts.length + this.invArtifacts.length;
    if (artifactCount === 0) {
      return html``;
    }

    return html`
      <milo-expandable-entry .hideContentRuler=${true}>
        <span slot="header">
          Artifacts: <span class="greyed-out">${artifactCount}</span>
        </span>
        <div slot="content">
          <ul>
            ${this.resultArtifacts.map((artifact) => html`
            <!-- TODO(weiweilin): refresh when the fetchUrl expires -->
            <li><a href=${artifact.fetchUrl} target="_blank">${artifact.artifactId}</a></li>
            `)}
          </ul>
          ${this.renderInvocationLevelArtifacts()}
        </div>
      </milo-expandable-entry>
    `;
  }

  private renderContent() {
    if (!this.shouldRenderContent) {
      return html``;
    }
    return html`
      ${this.renderSummaryHtml()}
      ${this.textDiffArtifact && html`
      <milo-text-diff-artifact .artifact=${this.textDiffArtifact}>
      </milo-text-diff-artifact>
      `}
      ${this.imageDiffArtifactGroup.diff && html`
      <milo-image-diff-artifact
        .expected=${this.imageDiffArtifactGroup.expected}
        .actual=${this.imageDiffArtifactGroup.actual}
        .diff=${this.imageDiffArtifactGroup.diff}
      >
      </milo-image-diff-artifact>
      `}
      ${this.renderArtifacts()}
      ${this.renderTags()}
    `;
  }

  protected render() {
    return html`
      <milo-expandable-entry
        .expanded=${this.expanded}
        .onToggle=${(expanded: boolean) => this.expanded = expanded}
      >
        <span id="header" slot="header">
          run #${this.id}
          <span class="${this.testResult.expected ? 'expected' : 'unexpected'}-result">
            ${this.testResult.expected ? '' : html`unexpectedly`}
            ${TEST_STATUS_DISPLAY_MAP[this.testResult.status]}
          </span>
          ${this.testResult.duration ? `after ${this.testResult.duration}` : ''}
        </span>
        <div slot="content">${this.renderContent()}</div>
      </milo-expandable-entry>
    `;
  }

  static styles = css`
    :host {
      display: block;
    }

    #header {
      font-size: 14px;
      letter-spacing: 0.1px;
      font-weight: 500;
    }

    .expected-result {
      color: var(--success-color);
    }
    .unexpected-result {
      color: var(--failure-color);
    }

    #summary-html {
      background-color: var(--block-background-color);
      padding: 5px;
    }
    #summary-html pre {
      margin: 0;
      font-size: 12px;
      white-space: pre-wrap;
    }

    .kv-key::after {
      content: ':';
    }
    .kv-value::after {
      content: ',';
    }
    .kv-value:last-child::after {
      content: '';
    }
    .greyed-out {
      color: var(--greyed-out-text-color);
    }

    ul {
      margin: 3px 0;
      padding-inline-start: 28px;
    }

    #inv-artifacts-header {
      margin-top: 12px;
    }
    `;
}

customElement('milo-result-entry')(
  consumeAppState(
    ResultEntryElement,
  ),
);
