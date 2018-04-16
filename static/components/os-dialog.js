let dialog = document.createElement('template');

dialog.innerHTML = `
    <style>
    :host {
        display: flex;
        width: 100%;
        height: 100%;
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

customElements.define('os-dialog', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(dialog.content.cloneNode(true))
    }
});