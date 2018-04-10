const location = document.createElement('template');

location.innerHTML = `
    <style>
    :host {
        all: initial;
      }
    div {
        display: flex;
        align-items: center;
        border: 1px solid gray;
        height: 22px;
        width: 200px;
    }
    div:focus-within {
        outline: 2px solid #006AFF;
        outline-offset: -1px;
    }
    input {
        border: none;
        height: 100%;
        width: 100%;
        padding: 0 4px;
    }
    input:focus {
        outline: none;
    }
    button {
        border: none;
        min-width: 22px;
        height: 100%
    }
    button:hover {
        background-color: gray;
    }
    button:focus {
        outline: none;
    }
    </style>
    <div>
        <input type="text" name="location"></input>
        <button>B</button>
        <slot></slot>
    </div>
`;

customElements.define('os-location-picker', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(location.content.cloneNode(true))
    }
});