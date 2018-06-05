let desktop = document.createElement('template');

desktop.innerHTML = `
    <style>
    :host {
        padding: 10px;
        width: 100%;
        height: 100%;
    }
    </style>
    <slot></slot>
`;

customElements.define('os-desktop', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(desktop.content.cloneNode(true))
    }
});