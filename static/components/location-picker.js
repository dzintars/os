const location = document.createElement('template');

location.innerHTML = `
    <style>
    div {
        --size: 22px;
        display: flex;
        align-items: center;
        border: thin solid gray;
        height: var(--size);
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
        width: var(--size);
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