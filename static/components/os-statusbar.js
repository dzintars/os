const statusbar = document.createElement('template');

statusbar.innerHTML = `
    <style>
    :host {
        display: flex;
        align-items: center;
        min-height: 22px;
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

customElements.define('os-statusbar', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(statusbar.content.cloneNode(true))
    }
});