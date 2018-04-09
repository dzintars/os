const taskbar = document.createElement('template');

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

customElements.define('os-taskbar', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(taskbar.content.cloneNode(true))
    }
});