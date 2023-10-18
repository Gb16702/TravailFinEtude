<script lang="ts">
    import Rotating from "./loaders/Rotating-circle.svelte";

    export let disabled: boolean = false;
    export let loading: boolean = false;
    let classes  = `text-white font-semibold rounded-[8px] transition duration-200 hover:bg-blue-500 md:w-fit px-6 ${disabled ? "bg-blue-200" : "bg-[#0B66EA]"}`
    export let additionalClasses: string = "";
</script>

<button type="submit" class={`${classes} ${additionalClasses}`} disabled={disabled}>
    {#if loading}
        <Rotating />
    {:else}
        <slot />
    {/if}
</button>