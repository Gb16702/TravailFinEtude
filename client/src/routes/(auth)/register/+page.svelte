<script lang="ts">
    import { enhance } from "$app/forms";
    import type { SubmitFunction } from "@sveltejs/kit";
    import type { Snapshot } from "./$types";

    let formData = {
        firstname: "",
        lastname: "",
        email: "",
        password: "",
        password_confirm: ""
    }

    export const snapshot: Snapshot = {
        capture: () => formData,
        restore: data => formData = data
    }

    let resultValue: any = null;
    const handleSubmit:SubmitFunction = () => {
        return async ({result, update}) => {
            resultValue = result;

            console.log(result);
            await update();
        }
    }

    const reduceStep = () => {
        return async ({result, update}: {result: any, update:any}) => {
            resultValue = result;
            await update();
        }
    }
</script>


{#if resultValue?.data?.step == 3}
    <div>
        <h2>Il ne vous reste plus qu'une étape</h2>
    </div>
{:else if resultValue?.data?.step == 2}
    <div class="bg-white flex items-center justify-center h-screen">
        <div class="min-w-[500px] min-h-[590px] bg-red-400 border border-black rounded-[6px]">
            <h2 class="w-full h-[70px] flex justify-center items-center text-[24px] font-semibold">Inscription à Host</h2>
            <form action="?/register" method="POST" use:enhance={handleSubmit}>
                <input type="hidden" name="step" value=2>
                <input bind:value={formData.password} type="password" name="password" placeholder="Mot de passe">
                <input bind:value={formData.password_confirm} type="password" name="password_confirm" placeholder="Confirmation de mot de passe">
                <button type="submit">Continuer</button>
            </form>
        </div>
    </div>
    {:else}
    <div class="bg-white flex items-center justify-center h-screen">
        <div class="min-w-[500px] min-h-[590px] bg-red-400 border border-black rounded-[6px]">
            <h2 class="w-full h-[70px] flex justify-center items-center text-[24px] font-semibold">Inscription à Host</h2>
            <form action="?/register" method="POST" use:enhance={handleSubmit}>
                <input type="hidden" name="step" value=1>
                <div>
                    <input bind:value={formData.firstname} type="text" name="firstname" placeholder="Votre prénom">
                    <input bind:value={formData.lastname} type="text" name="lastname" placeholder="Votre nom">
                </div>
                <input bind:value={formData.email} type="email" name="email" placeholder="Votre adresse mail">
                <button type="submit">Continuer</button>
            </form>
        </div>
    </div>
{/if}

<form action="?/step" method="POST" use:enhance={reduceStep}>
    <input type="hidden" value=resultValue>
    <button type="submit">Réduire</button>
</form>

