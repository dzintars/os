const taskbar = document.createElement('template');

taskbar.innerHTML = `
    <style>
    :host {
        display: flex;
        align-items: center;
        height: 36px;
        width: 100%;
    }
    :host([vbox]) {
        flex-direction: row;
    }
    :host([hbox]) {
        flex-direction: column;
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