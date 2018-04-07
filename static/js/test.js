// Create spans
let taskbar = document.createElement('template');
let viewport = document.createElement('template');

taskbar.innerHTML = `
    <style>
    :host {
        display: flex;
        align-items: center;
        border: 1px solid dodgerblue;
        height: 36px;
        width: 100%;
    }
    </style>
    <slot></slot>
`;

viewport.innerHTML = `
    <style>
    :host {
        display: flex;
        width: 100vw;
        height: 100vh;
        border: 1px solid #f00;
    }
    </style>
    <slot></slot>
`;

customElements.define('os-taskbar', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(taskbar.content.cloneNode(true))
    }
});

customElements.define('os-viewport', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(viewport.content.cloneNode(true))
    }
});