/**
 *  This function is used to detect a click outside of a component.
*
 * @param node
 * @returns
 */
const handleClickOutside = (node: HTMLElement): {destroy(): void} => {
    const click: (event: Event) => void = (event) => {
        if (node && !node.contains(event.target as Node) && !event.defaultPrevented) {
            node.dispatchEvent(new CustomEvent("click_outside"))
        }
    }
    document.addEventListener("click", click, true);
    return {
        destroy() {
            document.removeEventListener("click", click, true);
        }
    }
}

export default handleClickOutside;