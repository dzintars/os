let splitter = document.createElement('template');

splitter.innerHTML = `
    <style>
    :host {
        display: flex;
    }
    :host([hbox]) {
        flex-direction: row;
        width: 100%;
        height: 4px;
    }
    :host([vbox]) {
        flex-direction: column;
        height: 100%;
        width: 4px;
    }
    </style>
    <slot></slot>
`;

customElements.define('os-splitter', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(splitter.content.cloneNode(true))
    }
});