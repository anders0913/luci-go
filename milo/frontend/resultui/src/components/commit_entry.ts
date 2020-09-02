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
import { DateTime } from 'luxon';
import { computed, observable } from 'mobx';
import { DEFAULT_TIME_FORMAT } from '../libs/time_utils';

import { GitCommit } from '../services/milo_internal';
import './expandable_entry';

/**
 * Renders an expandable entry of the given commit.
 */
export class CommitEntryElement extends MobxLitElement {
  @observable.ref number = 0;
  @observable.ref repoUrl = '';
  @observable.ref commit!: GitCommit;
  onToggle = (_isExpanded: boolean) => {};

  @observable.ref private _expanded = false;
  get expanded() { return this._expanded; }
  set expanded(isExpanded) {
    if (isExpanded === this._expanded) {
      return;
    }
    this._expanded = isExpanded;
    this.onToggle(this._expanded);
  }

  @computed get commitTime() { return DateTime.fromISO(this.commit.committer.time); }
  @computed get title() { return this.commit.message.split('\n', 1)[0]; }
  @computed get descriptionLines() {
    const lines = this.commit.message.split('\n');
    lines.shift();
    if (lines[0].length === 0) {
      lines.shift();
    }
    return lines;
  }

  // TODO(crbug/1113995): support rendering links in the description.
  private renderDescriptionLine(line: string) { return html`${line}<br>`; }

  protected render() {
    return html`
      <milo-expandable-entry
        .expanded=${this.expanded}
        .onToggle=${(expanded: boolean) => this.expanded = expanded}
      >
        <span slot="header">
          <b>${this.number}. ${this.title}</b> <i>by ${this.commit.author.name} at ${this.commitTime.toFormat(DEFAULT_TIME_FORMAT)}</i>
        </span>
        <div slot="content">
          <table slot="content" border="0">
            <tr><td>Changed by:</td><td>${this.commit.author.name} - ${this.commit.author.email}</td></tr>
            <tr><td>Changed at:</td><td>${this.commitTime.toFormat(DEFAULT_TIME_FORMAT)}</td></tr>
            <tr><td>Revision:</td><td><a href=${`${this.repoUrl}/+/${this.commit.id}`} target="_blank">${this.commit.id}</a></td></tr>
          </table>
          <div id="summary">${this.descriptionLines.map((line) => this.renderDescriptionLine(line))}</div>
        </div>
      </milo-expandable-entry>
    `;
  }

  static styles = css`
    :host {
      display: block;
    }
    #summary {
      background-color: rgb(245, 245, 245);
      padding: 5px;
    }
    #summary > br:last-child {
      display: none;
    }
  `;
}

customElement('milo-commit-entry')(CommitEntryElement);