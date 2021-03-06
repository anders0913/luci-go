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
import { css, customElement, html, property } from 'lit-element';
import { computed, observable } from 'mobx';
import { fromPromise, IPromiseBasedObservable } from 'mobx-utils';
import { consumeContext } from '../../libs/context';
import { Artifact } from '../../services/resultdb';

/**
 * Renders a text artifact.
 */
export class TextArtifactElement extends MobxLitElement {
  @property({attribute: 'artifact-id'}) artifactID!: string;
  @observable.ref artifacts!: Map<string, Artifact>;

  @computed
  private get fetchUrl(): string|undefined {
    const artifact = this.artifacts.get(this.artifactID);
    return artifact ? artifact.fetchUrl : '';
  }

  @computed
  private get contentRes(): IPromiseBasedObservable<string> {
    if (!this.fetchUrl) {
      return fromPromise(Promise.race([]));
    }
    return fromPromise(fetch(this.fetchUrl).then((res) => res.text()));
  }

  @computed
  private get content() {
    return this.contentRes.state === 'fulfilled' ? this.contentRes.value : '';
  }

  protected render() {
    if (!this.content) {
      return html`
      <div id="load">
        Loading <milo-dot-spinner></milo-dot-spinner>
      </div>
      `;
    }
    return html`
        <pre>${this.content}</pre>
    `;
  }

  static styles = css`
    #load {
      color: var(--active-text-color);
    }
    pre {
      white-space: pre-wrap;
    }
  `;
}

customElement('text-artifact') (
  consumeContext<'artifacts', Map<string, Artifact>>('artifacts')(TextArtifactElement),
);
