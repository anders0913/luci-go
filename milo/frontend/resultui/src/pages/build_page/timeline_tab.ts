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
import { css, customElement, html } from 'lit-element';
import escape from 'lodash-es/escape';
import { autorun, observable } from 'mobx';
import { DataSet } from 'vis-data/peer';
import { Timeline } from 'vis-timeline/peer';

import { AppState, consumeAppState } from '../../context/app_state/app_state';
import { BuildState, consumeBuildState } from '../../context/build_state/build_state';
import { BUILD_STATUS_CLASS_MAP } from '../../libs/constants';
import { displayDuration } from '../../libs/time_utils';

export class TimelineTabElement extends MobxLitElement {
  @observable.ref appState!: AppState;
  @observable.ref buildState!: BuildState;
  @observable.ref rendered  = false;
  private disposers: Array<() => void> = [];

  connectedCallback() {
    super.connectedCallback();
    this.appState.selectedTabId = 'timeline';
    this.rendered = false;
    this.disposers.push(
      autorun(
        () => this.renderTimeline(),
      ),
    );
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    for (const disposer of this.disposers) {
      disposer();
    }
  }

  private renderTimeline() {
    if (this.buildState.build === null || !this.rendered) {
      return;
    }

    const container = this.shadowRoot!.getElementById('timeline')!;
    const options = {
      clickToUse: false,
      multiselect: false,
      orientation: {
        axis: 'both',
        item: 'top',
      },
      zoomable: false,
    };

    // Create a Timeline
    const timeline = new Timeline(
      container,
      new DataSet(this.buildState.build.steps.map((step, i) => ({
        className: `status ${BUILD_STATUS_CLASS_MAP[step.status]}`,
        content: '',
        id: i.toString(),
        group: i.toString(),
        start: (step.startTime || step.renderTime.get()).toMillis(),
        end: (step.endTime || step.renderTime.get()).toMillis(),
        type: 'range',
      }))),
      new DataSet(this.buildState.build.steps.map((step, i) => ({
        id: i.toString(),
        content: `
          <div class="group-title status ${BUILD_STATUS_CLASS_MAP[step.status]}">
            <span class="title">${escape(step.name)}</span>
            <span class="duration">( ${displayDuration(step.duration)} )</span>
          </div>
        `,
      }))),
      options,
    );

    timeline.on('select', (props) => {
      if (props.items.length === 0) {
        return;
      }
      const step = this.buildState.build!.steps[props.items[0]];
      const viewUrl = step.logs?.[0].viewUrl;
      if (!viewUrl) {
        return;
      }
      window.open(viewUrl, '_blank');
    });
  }

  protected render() {
    return html`
      <link rel="stylesheet" type="text/css" href="/static/styles/thirdparty/vis-timeline-graph2d.min.css">
      <div id="timeline"></div>
    `;
  }

  protected firstUpdated() {
    this.rendered = true;
  }

  static styles = css`
    .vis-range {
      cursor: pointer;
    }

    .status.infra-failure {
      color: #FFFFFF;
      background-color: #c6c;
      border-color: #ACA0B3;
    }
    .status.started {
      color: #000;
      background-color: #fd3;
      border-color: #C5C56D;
    }
    .status.failure {
      color: #000;
      background-color: #e88;
      border-color: #A77272;
      border-style: solid;
    }
    .status.canceled {
      color: #000;
      background-color: #8ef;
      border-color: #00d8fc;
      border-style: solid;
    }
    .status.success {
      color: #000;
      background-color: #8d4;
      border-color: #4F8530;
    }

    .group-title > .title{
      display: inline-block;
      white-space: nowrap;
      max-width: 50em;
      overflow: hidden;
      text-overflow: ellipsis;
    }

    .group-title > .duration{
      display: inline-block;
      text-align: right;
      float: right;
    }

    .group-title {
      font-weight: bold;
      padding: 5px;
    }

    .vis-labelset .vis-label .vis-inner {
      display: block;
    }
  `;
}

customElement('milo-timeline-tab')(
  consumeBuildState(
    consumeAppState(TimelineTabElement),
  ),
);
