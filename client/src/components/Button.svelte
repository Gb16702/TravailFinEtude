<script lang="ts">
    import Rotating from "./loaders/Rotating-circle.svelte";

    let classes  = "text-white font-semibold text-[16px] bg-accent-violet-d rounded-[5px] h-[40px] transition duration-200 hover:bg-accent-violet-c/[.85]"

    export let disabled: boolean = false;
    export let loading: boolean = false;
    export let additionalClasses: string = "";
</script>

<button class={`${classes} ${additionalClasses}`} disabled={disabled}>
    {#if loading}
        <Rotating />
    {:else}
        <slot />
    {/if}
</button>