let container = document.createElement('template');

container.innerHTML = `
    <style>
    :host {
        display: flex;
        width: 100%;
        height: 100%;
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

customElements.define('os-container', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(container.content.cloneNode(true))
    }
});