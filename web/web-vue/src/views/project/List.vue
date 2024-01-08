<script setup>
import { ref, onMounted } from 'vue'
import { projectList } from '@/api/project.js'

const dataList = ref(null)

onMounted(() => {
  loadProjectList()
})

const loadProjectList = () => {
  projectList({}).then(res => {
    console.log(res)
  })
}
</script>

<template>
  <div class="grid">
    <div class="col-12">
      <div class="card">
        <DataTable
          ref="dt"
          :value="dataList"
        >
        </DataTable>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>