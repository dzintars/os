const panel = document.createElement('template');

panel.innerHTML = `
    <style>
    :host {
        display: flex;
        flex-grow: 1;
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

customElements.define('os-panel', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(panel.content.cloneNode(true))
    }
});