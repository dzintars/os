let row = document.createElement('template');

row.innerHTML = `
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

customElements.define('os-row', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(row.content.cloneNode(true))
    }
});