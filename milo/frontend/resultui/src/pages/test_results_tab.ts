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
import '@material/mwc-icon';
import { css, customElement, html } from 'lit-element';
import { repeat } from 'lit-html/directives/repeat';
import { styleMap } from 'lit-html/directives/style-map';
import { computed, observable, reaction } from 'mobx';

import '../components/left_panel';
import '../components/test_filter';
import { TestFilter } from '../components/test_filter';
import '../components/test_nav_tree';
import '../components/variant_entry';
import { VariantEntryElement } from '../components/variant_entry';
import { AppState, consumeAppState } from '../context/app_state/app_state';
import { consumeInvocationState, InvocationState } from '../context/invocation_state/invocation_state';
import { ReadonlyVariant, TestNode, VariantStatus } from '../models/test_node';

/**
 * Display a list of test results.
 */
export class TestResultsTabElement extends MobxLitElement {
  @observable.ref appState!: AppState;
  @observable.ref invocationState!: InvocationState;

  private disposers: Array<() => void> = [];
  private async loadNextPage() {
    try {
      await this.invocationState.testLoader.loadNextPage();
    } catch (e) {
      this.dispatchEvent(new ErrorEvent('error', {
        error: e,
        message: e.toString(),
        composed: true,
        bubbles: true,
      }));
    }
  }

  @computed
  private get hasSingleVariant() {
    // this operation should be fast since the iterator is executed only when
    // there's only one test.
    return this.invocationState.selectedNode.testCount === 1 && [...this.invocationState.selectedNode.tests()].length === 1;
  }

  private toggleAllVariants(expand: boolean) {
    this.shadowRoot!.querySelectorAll<VariantEntryElement>('milo-variant-entry')
      .forEach((e) => e.expanded = expand);
  }

  connectedCallback() {
    super.connectedCallback();
    this.appState.selectedTabId = 'test-results';

    // When a new test loader is received, load the first page and reset the
    // selected node.
    this.disposers.push(reaction(
      () => this.invocationState.testLoader,
      (testLoader) => {
        this.loadNextPage();
        this.invocationState.selectedNode = testLoader.node;
      },
      {fireImmediately: true},
    ));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    for (const disposer of this.disposers) {
      disposer();
    }
  }

  private renderAllVariants() {
    const exoneratedVariants: ReadonlyVariant[] = [];
    const expectedVariants: ReadonlyVariant[] = [];
    const unexpectedVariants: ReadonlyVariant[] = [];
    const flakyVariants: ReadonlyVariant[] = [];
    for (const test of this.invocationState.selectedNode.tests()) {
      for (const variant of test.variants) {
        switch (variant.status) {
          case VariantStatus.Exonerated:
            exoneratedVariants.push(variant);
            break;
          case VariantStatus.Expected:
            expectedVariants.push(variant);
            break;
          case VariantStatus.Unexpected:
            unexpectedVariants.push(variant);
            break;
          case VariantStatus.Flaky:
            flakyVariants.push(variant);
            break;
          default:
            console.error('unexpected variant type', variant);
            break;
        }
      }
    }
    return html`
      ${unexpectedVariants.length === 0 ? html`
      <div class="list-entry">No unexpected test results.</div>
      <hr class="divider">
      ` : ''}
      ${this.renderVariants(unexpectedVariants)}
      ${this.renderVariants(flakyVariants)}
      ${this.renderVariants(exoneratedVariants)}
      ${this.renderVariants(expectedVariants)}
    `;
  }

  private renderVariants(variants: ReadonlyVariant[]) {
    return html`
      ${repeat(
        variants.map((v, i, variants) => [variants[i-1], v, variants[i+1]] as [ReadonlyVariant | undefined, ReadonlyVariant, ReadonlyVariant | undefined]),
        ([_, v]) => `${v.testId} ${v.variantHash}`,
        ([prev, v, next]) => html`
        <milo-variant-entry
          .variant=${v}
          .prevTestId=${prev?.testId ?? ''}
          .prevVariant=${prev?.testId === v.testId ? prev : null}
          .expanded=${this.hasSingleVariant}
          .displayVariantId=${prev?.testId === v.testId || next?.testId === v.testId}
        ></milo-variant-entry>
      `)}
      ${variants.length !== 0 ? html`<hr class="divider">` : ''}
    `;
  }

  protected render() {
    const state = this.invocationState;

    return html`
      <div id="header">
        <milo-test-filter
          .onFilterChanged=${(filter: TestFilter) => {
            this.invocationState.showExonerated = filter.showExonerated;
            this.invocationState.showExpected = filter.showExpected;
            this.invocationState.showFlaky = filter.showFlaky;
          }}
        >
        </milo-test-filter>
        <mwc-button
          class="action-button"
          dense unelevated
          @click=${() => this.toggleAllVariants(true)}
        >Expand All</mwc-button>
        <mwc-button
          class="action-button"
          dense unelevated
          @click=${() => this.toggleAllVariants(false)}
        >Collapse All</mwc-button>
      </div>
      <div id="main">
        <milo-left-panel>
          <milo-test-nav-tree
            .testLoader=${state.testLoader}
            .onSelectedNodeChanged=${(node: TestNode) => state.selectedNode = node}
          ></milo-test-nav-tree>
        </milo-left-panel>
        <div id="test-result-view">
          ${this.renderAllVariants()}
          <div class="list-entry">
            <span>Showing ${state.selectedNode.testCount} tests.</span>
            <span
              id="load"
              style=${styleMap({'display': state.testLoader.done ? 'none' : ''})}
            >
              <span
                id="load-more"
                style=${styleMap({'display': state.testLoader.isLoading ? 'none' : ''})}
                @click=${this.loadNextPage}
              >
                Load More
              </span>
              <span
                style=${styleMap({'display': state.testLoader.isLoading ? '' : 'none'})}
              >
                Loading <milo-dot-spinner></milo-dot-spinner>
              </span>
              <mwc-icon id="load-info" title="Newly loaded entries might be inserted into the list.">info</mwc-icon>
            </span>
          </div>
        </div>
      </div>
    `;
  }

  static styles = css`
    :host {
      display: grid;
      grid-template-rows: auto 1fr;
      overflow-y: hidden;
    }

    #header {
      display: grid;
      grid-template-columns: 1fr auto auto;
      grid-gap: 5px;
      height: 28px;
      padding: 5px 10px 3px 10px;
    }
    milo-test-filter {
      margin: 5px;
      margin-bottom: 0px;
    }
    .action-button {
      --mdc-theme-primary: rgb(0, 123, 255);
    }

    #main {
      display: flex;
      border-top: 1px solid #DDDDDD;
      overflow-y: hidden;
    }
    #test-result-view {
      flex: 1;
      overflow-y: auto;
      padding-top: 5px;
    }
    #test-result-view>* {
      margin-bottom: 2px;
    }

    .divider {
      border: none;
      border-top: 1px solid #DDDDDD;
    }

    milo-test-nav-tree {
      overflow: hidden;
    }

    .list-entry {
      margin-left: 5px;
      margin-top: 5px;
    }
    #load {
      color: blue;
    }
    #load-more {
      color: blue;
      cursor: pointer;
    }
    #load-info {
      color: #212121;
      --mdc-icon-size: 1.2em;
      vertical-align: bottom;
    }
  `;
}

customElement('milo-test-results-tab')(
  consumeInvocationState(
      consumeAppState(TestResultsTabElement),
  ),
);
