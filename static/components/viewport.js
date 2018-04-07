let viewport = document.createElement('template');

viewport.innerHTML = `
    <style>
    :host {
        display: flex;
        width: 100vw;
        height: 100vh;
        border: 1px solid #f00;
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