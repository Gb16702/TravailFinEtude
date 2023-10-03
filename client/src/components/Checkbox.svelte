<script lang="ts">
    export let theme: "default" | "alt" = "default";
    export let checkedClasses: string = "";
    export let checked: boolean;

    let defaultClasses: string = "absolute -translate-x-1/2 -translate-y-1/2 pointer-events-none top-1/2 left-1/2 w-full h-full flex items-center justify-center overflow-hidden rounded-[4px]"

    if(theme === "default") {
        checkedClasses="bg-black"
    } else if (theme === "alt") {
        checkedClasses="bg-accent-light-mode-a"
    }

    export let onCheck: () => void = () => {};
</script>
<div class="relative w-fit h-fit overflow-hidden border rounded-[50px]">
    <div class="appearance-none rounded-[4px] bg-white overflow-hidden w-5 h-5" on:click={onCheck}></div>
    {#if checked}
    <div class={`${defaultClasses} ${checkedClasses}`}>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="3" stroke="white" class="w-[70%] h-[70%]">
        <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 12.75l6 6 9-13.5" />
    </svg>
</div>
{/if}
</div>

<style>
    .checkbox{
        appearance: none;
        border-radius: 50px;
        background-color: white;
        overflow: hidden;
    }
</style>