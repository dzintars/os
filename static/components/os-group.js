let group = document.createElement('template');

group.innerHTML = `
    <style>
    :host {
        display: flex;
    }
    :host([hbox]) {
        flex-direction: row;
    }
    :host([vbox]) {
        flex-direction: column;
        justify-content: center;
    }
    </style>
    <slot></slot>
`;

customElements.define('os-group', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(group.content.cloneNode(true))
    }
});