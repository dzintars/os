let stepbar = document.createElement('template');

stepbar.innerHTML = `
    <style>
    :host {
        display: flex;
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

customElements.define('os-stepbar', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(stepbar.content.cloneNode(true))
    }
});