let row = document.createElement('template');

row.innerHTML = `
    <style>
    :host {
        display: flex;
        width: 100%;
    }
    :host([vbox]) {
        flex-direction: row;
    }
    :host([hbox]) {
        flex-direction: column;
    }
    :host([center]) {
        justify-content: center;
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