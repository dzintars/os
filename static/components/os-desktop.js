let desktop = document.createElement('template');

desktop.innerHTML = `
    <style>
    :host {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(48px, 1fr));
        grid-template-rows: repeat(auto-fill, minmax(48px, 1fr));
        padding: 10px;
        grid-gap: 10px;
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

customElements.define('os-desktop', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(desktop.content.cloneNode(true))
    }
});