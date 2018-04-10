let viewport = document.createElement('template');

viewport.innerHTML = `
    <style>
    :host {
        display: flex;
        width: 100vw;
        height: 100vh;
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

customElements.define('os-viewport', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(viewport.content.cloneNode(true))
    }
});