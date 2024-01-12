<script setup>
import Toast from 'primevue/toast'
</script>

<template>
  <Toast group="network">
    <template #message="slotProps">
      <div flex flex-col flex-grow-1>
        <div p-grid>
          <div p-col-12 font-500 mb-5px>{{ slotProps.message.summary }}</div>
        </div>
        <div p-grid flex>
          <div p-col-2 h-24px>
            <Tag :value="slotProps.message.method"></Tag>
          </div>
          <div p-col-10 line-height-24px ml-5px>
            {{ slotProps.message.detail }}
          </div>
        </div>
        <InlineMessage v-show="slotProps.message.msg" mt-5px severity="warn">{{ slotProps.message.msg }}</InlineMessage>
      </div>
    </template>
  </Toast>
  <router-view></router-view>
</template>

<style scoped>
</style>
