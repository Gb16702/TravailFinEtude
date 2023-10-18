<script lang="ts">
    import {page} from "$app/stores";
    import { goto } from "$app/navigation";
    import ModernArrow from "./icons/ModernArrow.svelte";

    let baseClasses: string = "bg-white outline outline-1 outline-black h-[40px] rounded-lg text-white font-semibold flex items-center justify-center px-3 gap-x-[6px]"
    export let additionalClasses: string | null;

    let currentURL: string[] = [
        "/auth/register",
        "/auth/login",
        "/auth"
    ]

    let newURL: string | undefined = currentURL.find((url) => url === $page.url.pathname)
    ?.split("/").slice(0, -1).join("/").replace(/^$/, "/")

    function handleClick() {
        goto(newURL as string);
    }
</script>

<button on:click={handleClick} class={`${baseClasses} ${additionalClasses}`}>
    <ModernArrow classes={"fill-black w-4 h-4 relative right-[2px]"} />
    <h2 class="text-black font-bold text-[15px] uppercase tracking-[0.015rem]">Retour</h2>
</button>