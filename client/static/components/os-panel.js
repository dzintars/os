const panel = document.createElement('template');

panel.innerHTML = `
    <style>
    :host {
        display: flex;
        flex: 1;
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

customElements.define('os-panel', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(panel.content.cloneNode(true))
    }
});