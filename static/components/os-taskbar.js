const taskbar = document.createElement('template');

taskbar.innerHTML = `
    <style>
    :host {
        display: flex;
        align-items: center;
        height: 36px;
        width: 100%;
    }
    :host([hbox]) {
        flex-direction: row;
    }
    :host([vbox]) {
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