const shortcutbar = document.createElement('template');

shortcutbar.innerHTML = `
    <style>
    :host {
        display: flex;
    }
    :host([hbox]) {
        flex-direction: row;
        align-items: center;
    }
    :host([vbox]) {
        flex-direction: column;
        justify-items: center;
    }
    </style>
    <slot></slot>
`;

customElements.define('os-shortcutbar', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(shortcutbar.content.cloneNode(true))
    }
});