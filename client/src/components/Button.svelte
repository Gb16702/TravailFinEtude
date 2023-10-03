<script lang="ts">
    import Rotating from "./loaders/Rotating-circle.svelte";

    let classes  = "text-white font-semibold  bg-black rounded-[6px] h-[40px] transition duration-200 hover:bg-zinc-900/[.85] md:w-fit px-6"

    export let disabled: boolean = false;
    export let loading: boolean = false;
    export let additionalClasses: string = "";
</script>

<button type="submit" class={`${classes} ${additionalClasses}`} disabled={disabled}>
    {#if loading}
        <Rotating />
    {:else}
        <slot />
    {/if}
</button>