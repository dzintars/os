const toolbar = document.createElement('template');

toolbar.innerHTML = `
    <style>
    :host {
        display: flex;
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

customElements.define('os-toolbar', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(toolbar.content.cloneNode(true))
    }
});