<script lang="ts">
    type Input = {
        input_type: "text" | "password" | "email" | "number" | "tel" | null;
        name: string;
        placeholder: string;
        label?: string | null;
        value: string;
        additionalClasses: string;
        tag?: string;
        passwordClasses?: string;
        IsPasswordInput: boolean;
    }
    let inputBaseClass="w-full h-full px-2 overflow-hidden placeholder-zinc-400/[.55] outline-none text-[15.2px] bg-transparent"
    let classes = "rounded-[5px] overflow-hidden h-[40px] transition-all bg-zinc-100/[.8] duration-200 focus-within:outline-accent-lilas px-2  outline outline-1 outline-zinc-200/[.4]"

    export let type: Input["input_type"] = "text"
    export let name: Input["name"] = ""
    export let placeholder: Input["placeholder"] = ""
    export let label: Input["label"] = null
    export let value: Input["value"] = ""
    export let additionalClasses: Input["additionalClasses"] = ""
    export let passwordClasses: Input["passwordClasses"] = ""
</script>

<div class={`${label && "flex flex-col gap-y-1"} ${additionalClasses}`}>
    {#if label}
        <label for={name} class="font-semibold text-[14.4px]">{label}</label>
    {/if}
    <div class={`${passwordClasses} ${classes}`}>
        {#if type === "text" }
            <input
                placeholder={placeholder}
                bind:value={value}
                type="text"
                class="w-full h-full px-2 overflow-hidden placeholder-zinc-300/[.9] outline-none text-[15.2px] bg-transparent"
            />
            <slot />
        {:else if type==="email"}
            <input
                class="w-full h-full px-2 overflow-hidden placeholder-zinc-300/[.9] outline-none text-[15.2px] bg-transparent"
                bind:value={value}
                placeholder={placeholder}
                type="email"
            />
        {:else if type==="password"}
            <input
                class={inputBaseClass}
                bind:value={value}
                placeholder={placeholder}
                type="password"
            />
            <slot />
        {:else if type==="tel"}
            <input
                placeholder={placeholder}
                bind:value={value}
                type="tel"
                class="w-full h-full px-2 overflow-hidden placeholder-zinc-300/[.9] outline-none text-[15.2px] bg-transparent"
            />
        {/if}
    </div>
</div>
