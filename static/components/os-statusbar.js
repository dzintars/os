const statusbar = document.createElement('template');

statusbar.innerHTML = `
    <style>
    :host {
        display: flex;
        align-items: center;
        min-height: 22px;
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

customElements.define('os-statusbar', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(statusbar.content.cloneNode(true))
    }
});