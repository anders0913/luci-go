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
import { css, customElement } from 'lit-element';
import { html } from 'lit-html';
import { styleMap } from 'lit-html/directives/style-map';
import { DateTime } from 'luxon';
import { computed, observable } from 'mobx';
import { BUILD_STATUS_CLASS_MAP, BUILD_STATUS_DISPLAY_MAP, BUILD_STATUS_ICON_MAP } from '../libs/constants';

import { displayDuration } from '../libs/time_utils';
import { ChainableURL, renderMarkdown } from '../libs/utils';
import { BuildStatus } from '../services/buildbucket';
import { StepExt } from '../services/build_page';
import './expandable_entry';

/**
 * Renders a step.
 */
@customElement('milo-build-step-entry')
export class BuildStepEntryElement extends MobxLitElement {
  @observable.ref number = 0;
  @observable.ref expanded = false;
  @observable.ref step!: StepExt;
  @observable.ref showDebugLogs = false;

  toggleAllSteps(expand: boolean) {
    this.expanded = expand;
    this.shadowRoot!.querySelectorAll<BuildStepEntryElement>('milo-build-step-entry')
      .forEach((e) => e.toggleAllSteps(expand));
  }

  @computed private get shortName() { return this.step.name.split('|')[0] || 'ERROR: Empty Name'; }

  @computed private get duration() {
    const start = DateTime.fromISO(this.step.interval.start);
    const end = DateTime.fromISO(this.step.interval.end);
    const now = DateTime.fromISO(this.step.interval.now);
    return this.step.status === BuildStatus.Started ? now.diff(start) : end.diff(start);
  }

  @computed private get header() {
    return this.step.summary_markdown?.split(/\<br\/?\>/i)[0] || '';
  }

  @computed private get summary() {
    return this.step.summary_markdown?.split(/\<br\/?\>/i).slice(1).join('<br>') || '';
  }

  @computed private get logs() {
    const logs = this.step.logs || [];
    return this.showDebugLogs ? logs : logs.filter((log) => !log.name.startsWith('$'));
  }

  protected render() {
    return html`
      <milo-expandable-entry .expanded=${this.expanded}>
        <span slot="header">
          <mwc-icon
            id="status-indicator"
            class=${BUILD_STATUS_CLASS_MAP[this.step.status]}
            title=${BUILD_STATUS_DISPLAY_MAP[this.step.status]}
          >${BUILD_STATUS_ICON_MAP[this.step.status]}</mwc-icon>
          <b>${this.number}. ${this.shortName}</b>
          <span id="header-markdown">${renderMarkdown(this.header)}</span>
          <span id="duration">${displayDuration(this.duration)}</span>
        </span>
        <div slot="content">
          <div id="summary" style=${styleMap({display: this.summary ? '' : 'none'})}>
            ${renderMarkdown(this.summary)}
          </div>
          <ul id="log-links" style=${styleMap({display: this.step.logs?.length ? '' : 'none'})}>
            ${this.logs.map((log) => html`
            <li>
              <a href=${log.view_url} target="_blank">${log.name}</a>
              <a
                style=${styleMap({'display': ['stdout', 'stderr'].indexOf(log.name) !== -1 ? '' : 'none'})}
                href=${new ChainableURL(log.view_url).withSearchParam('format', 'raw', true).toString()}
                target="_blank"
              >[raw]</a>
            </li>
            `)}
          </ul>
          ${this.step.children?.map((child, i) => html`
          <milo-build-step-entry
            class="list-entry"
            .number=${i + 1}
            .step=${child}
            .showDebugLogs=${this.showDebugLogs}
          ></milo-build-step-entry>
          `) || ''}
        </div>
      </milo-expandable-entry>
    `;
  }

  static styles = css`
    :host {
      display: block;
    }

    #status-indicator {
      vertical-align: bottom;
    }
    #status-indicator.started {
      color: var(--started-color);
    }
    #status-indicator.success {
      color: var(--success-color);
    }
    #status-indicator.failure {
      color: var(--failure-color);
    }
    #status-indicator.infra-failure {
      color: var(--critical-failure-color);
    }
    #status-indicator.canceled {
      color: var(--canceled-color);
    }

    #header-markdown > p {
      display: inline;
    }

    #duration {
      color: white;
      background-color: var(--active-color);
      display: inline-block;
      padding: .25em .4em;
      font-size: 75%;
      font-weight: 700;
      line-height: 1;
      text-align: center;
      white-space: nowrap;
      vertical-align: baseline;
      border-radius: .25rem;
    }

    #summary {
      background-color: var(--block-background-color);
      padding: 5px;
    }

    #summary > p:first-child {
      margin-block-start: 0px;
    }

    #summary > p:last-child {
      margin-block-end: 0px;
    }

    #log-links {
      margin: 3px 0;
      padding-inline-start: 28px;
    }

    #log-links>li:nth-child(odd) {
      background-color: var(--block-background-color);
      list-style-type: circle;
    }

    #log-links>li {
      list-style-type: circle;
      padding: 0.1em 1em 0.1em 1em;
    }
    #log-links>li>a {
      color: var(--default-text-color);
    }

    milo-build-step-entry {
      margin-bottom: 2px;
    }
  `;
}