<template>
    <button @click="buildImg" :aria-label="$t('buttons.buildImg')" :title="$t('buttons.buildImg')" class="action" id="build-img-button">
        <i class="material-icons">donut_small</i>
        <span>{{ $t('buttons.buildImg') }}</span>
    </button>
</template>

<script>
import { files as api } from '@/api'
export default {
    name: 'build-img-button',
    methods: {
        buildImg: function () {
            api.buildImg()
        }
    }
}
</script>
