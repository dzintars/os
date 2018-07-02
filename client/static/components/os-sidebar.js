const sidebar = document.createElement('template');

sidebar.innerHTML = `
    <style>
    :host {
        display: flex;
        width: 240px;
        min-width: 240px;
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

customElements.define('os-sidebar', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(sidebar.content.cloneNode(true))
    }
});