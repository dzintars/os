const tabbar = document.createElement('template');

tabbar.innerHTML = `
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

customElements.define('os-tabbar', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(tabbar.content.cloneNode(true))
    }
});