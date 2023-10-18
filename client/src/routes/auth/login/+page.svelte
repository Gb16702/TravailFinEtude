<script lang="ts">
	import { enhance } from '$app/forms';
    import type { Snapshot } from "@sveltejs/kit";
    import type { ActionData, PageData, SubmitFunction } from './$types';
    import Input from '../../../components/Input.svelte';
    import Button from '../../../components/Button.svelte';
    import Menu from '../../../components/Menu.svelte';
  import AuthSectionToggler from '../../../components/AuthSectionToggler.svelte';

    let loginFormDataObject:loginFormDataObject = {
        email: {
            type: "email",
            name:"email",
            label: "Adresse mail",
            value: "",
        },
        password: {
            type: "password",
            name:"password",
            label: "Mot de passe",
            value: "",
        }
    }

    export const snapshot: Snapshot = {
        capture: () => loginFormDataObject,
        restore: data => loginFormDataObject = data
    }


    function handleSubmit(){
        return async ({update}: any) => {
            await update();
        }
    }

    export let form: ActionData;

    console.log(form);

</script>

<Menu />
<div class="h-[15px]"></div>
<div class="w-full flex items-start justify-center flex-col gap-y-[2px] h-[30%]">
        <div class="flex items-center justify-center flex-col w-full h-[150px]">
            <h1 class="text-[19px] text-black font-medium">Bienvenue</h1>
            <h1 class="font-bold text-[31px] text-[#383838]">Connectez-vous ici</h1>
        </div>
</div>
<div class="w-full flex flex-col items-center justify-center min-h-[480px]">
    <div class="w-[97%] md:w-[75%] lg:w-[650px] px-2">
        <form action="?/login" method="POST" use:enhance={handleSubmit}>
            <Input type={loginFormDataObject.email.type} name={loginFormDataObject.email.name} label={loginFormDataObject.email.label} bind:value={loginFormDataObject.email.value} additionalClasses="w-full"  />
            {#if form?.errors?.email}
                <p class="text-red-500 text-[13px] mt-1">{form?.errors?.email}</p>
            {/if}
            <Input type={loginFormDataObject.password.type} name={loginFormDataObject.password.name} label={loginFormDataObject.password.label} bind:value={loginFormDataObject.password.value} additionalClasses="w-full"  />
            {#if form?.errors?.password}
                <p class="text-red-500 text-[13px] mt-1">{form?.errors?.password}</p>
            {/if}
            <div class="w-full md:flex md:justify-between md:items-center md:bg-[#EDEEF2] mt-4 md:px-2 md:rounded-[8px] md:py-3">
                <div class="max-md:hidden">
                    <AuthSectionToggler />
                </div>
                <Button additionalClasses="w-full h-[40px] md:h-[36px]">
                    Se connecter
                </Button>
            </div>
        </form>
    </div>
    <div class="mt-6 md:hidden">
        <AuthSectionToggler />
    </div>
</div>
