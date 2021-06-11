<template>
  <div class="goal">
    <div>
      <button @click="remove">Eliminar</button>
      <div v-if="showNested">
        <div v-for="(task, index) in nestedTask" :key="task.id">
          </NestedTask />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Goal'
}
</script>
