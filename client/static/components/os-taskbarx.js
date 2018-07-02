const taskbarx = document.createElement('template');

taskbarx.innerHTML = `
<link rel="import" href="/static/img/iron-icons/maps-icons.html">
    <style>
    :host {
        display: flex;
        flex-direction: row;
        align-items: center;
        width: 100%;
        height: 37px;
        min-height: 37px;
        background: #FFFFFF;
        border-bottom: 1px solid #E6E6E6;
    }
    .open-launcher {
        display: grid;
        place-items: center;
        border: none;
        min-width: 37px;
        height: 100%;
        background-color: dodgerblue;
        cursor: pointer;
        margin-right: 8px;
        & svg {
            fill: white;
        }
      }
    .i-icon {
        height: 22px;
        width: 22px;
        cursor: context-menu;
        stroke-width: 2;
    }
    .app-control-icon {
        height: 37px;
        width: 37px;
        min-width: 37px;
        display: grid;
        place-items: center;
        background: lightgray;
        cursor: context-menu;
        &.danger {
            background: red;
            & .i-icon {
                color: white;
            }
        }
        & .i-icon {
            // ToDo: Tooltip on hover appears only on icon and not on the whole button
            color: white;
            cursor: context-menu;
        }
    }
    .actions {
        align-items: center;
        &-icon {
            display: grid;
            place-items: center;
            width: 28px;
            height: 28px;
            border-radius: 3px;
            background: $light10;
            margin-left: 4px;
            cursor: pointer;
            &:last-child {
                margin-right: 8px;
            }
            & .i-icon {
                color: $blue;
            }
            &:hover {
                background: $light20;
            }
        }
    }
    </style>
    <button id="btn-open-launcher" class="open-launcher">
        <img class="i-icon" src="/static/img/svg-icons/9dots.svg" alt="Kiwi standing on oval">
    </button>
    <slot name="path"></slot>
    <os-panel hbox></os-panel>
    <os-group hbox class="actions">
        <div class="actions-icon">
            <svg class="i-icon"><use xlink:href="/static/img/svg-icons/icons.svg#i-search"></use></svg>
        </div>
        <div class="actions-icon">
            <svg class="i-icon"><use xlink:href="/static/img/svg-icons/icons.svg#i-bell"></use></svg>
        </div>
        <div class="actions-icon">
            <svg class="i-icon"><use xlink:href="/static/img/svg-icons/icons.svg#i-info"></use></svg>
        </div>
        <iron-icon icon="maps:directions-bus"></iron-icon>
        <div class="app-control-icon" onclick="toggleFullScreen()">
            <img class="i-icon" src="/static/img/svg-icons/9dots.svg" alt="Kiwi standing on oval">
        </div>
        <a href="/apps/crm/customers" title="Exit Application">
            <div class="app-control-icon danger">
                <img class="i-icon" src="/static/img/svg-icons/9dots.svg" alt="Kiwi standing on oval">
            </div>
        </a>
    </os-group>
`;

customElements.define('os-taskbarx', class extends HTMLElement {
    constructor() {
        super();
        const shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(taskbarx.content.cloneNode(true))
    }
});