<script lang="ts">
    type Arrow = {
        rotation: null | "90" | "180" | "270";
        classes: string;
        stroke_width: string;
    }

    export let rotation: Arrow["rotation"] = null;
    export let classes: Arrow["classes"] = "w-5 h-5";
    export let stroke_width: Arrow["stroke_width"] = "1.5";

    let rotationClasses: string = `rotate-[${rotation}deg]`
</script>

<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width={stroke_width} class={`${classes} ${rotationClasses}`}>
    <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
</svg>
