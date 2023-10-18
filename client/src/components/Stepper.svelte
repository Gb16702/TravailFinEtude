<script lang="ts">
    export let step: number;
</script>

<div class="flex flex-row w-full items-center gap-x-[9px] justify-center h-fit py-4">
    <h3 class="text-zinc-400 text-center text-[14px]" class:active={step >= 1}>Informations</h3>
    <div class="w-[15px] h-[1px] bg-zinc-400/[.4] relative top-[1px]"></div>
    <h3 class="text-zinc-400/[.4] text-center text-[14px]" class:active={step >= 2}>Sécurité</h3>
    <div class="w-[15px] h-[1px] bg-zinc-400/[.4] relative top-[1px]"></div>
    <h3 class="text-zinc-400/[.4] text-center text-[14px]" class:active={step == 3}>Validation</h3>
</div>

<style>
    .active {
        color: black;
        font-weight: 500;
    }
</style>