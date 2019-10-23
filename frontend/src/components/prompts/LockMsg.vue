<template>
    <div class="card floating">
        <div class="card-title">
            <h2>{{ $t('lock.title') }}</h2>
        </div>

        <div class="card-content">
            <p>{{ $t('lock.msg') }}</p>
        </div>

        <div class="card-action">
            <button type="submit"
                    @click="$store.commit('closeHovers')"
                    class="button button--flat"
                    :aria-label="$t('buttons.ok')"
                    :title="$t('buttons.ok')">{{ $t('buttons.ok') }}</button>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'lock-msg',
    }
</script>
