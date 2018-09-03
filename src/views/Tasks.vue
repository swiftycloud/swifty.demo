<template>
  <div class="tasks">
    <el-row>
      <el-col>
        <el-tabs v-model="activeTabName">
          <el-tab-pane label="TASKS" name="tasks"></el-tab-pane>
          <el-tab-pane label="PROFILE" name="profile"></el-tab-pane>
        </el-tabs>
      </el-col>
    </el-row>

    <el-row>
      <el-col :lg="10">
        <div class="tasks-filter">
          <el-button size="mini" @click="filter = 'all'" :class="{ active: filter === 'all' }">All</el-button>
          <el-button size="mini" @click="filter = 'done'" :class="{ active: filter === 'done' }">Completed</el-button>
          <el-button size="mini" @click="filter = 'new'" :class="{ active: filter === 'new' }">Active</el-button>
        </div>

        <div class="tasks-form">
          <input type="text" @keyup.enter="addTask" v-model="task" placeholder="Enter your task here" class="tasks-input">
          <a href="#" class="tasks-button" @click.prevent.stop="addTask">
            <i class="el-icon-circle-plus-outline"></i>
          </a>
        </div>

        <ul class="tasks-list" v-if="tasks.length">
          <li class="tasks-item" v-for="task in tasks" :key="task.id">
            <a href="#" class="tasks-item__check" @click.prevent.stop="doneTask(task)" v-if="task.status !== 'done'">
              <i class="el-icon-circle-check-outline"></i>
            </a>
            <span class="tasks-item__check checked" v-if="task.status === 'done'">
              <i class="el-icon-circle-check-outline"></i>
            </span>
            <span class="tasks-item__name">{{ task.task }}</span>
            <a href="#" class="tasks-item-delete" @click.prevent.stop="deleteTask(task)">
              <i class="el-icon-delete"></i>
            </a>
          </li>
        </ul>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  data () {
    return {
      activeTabName: 'tasks',
      task: null,
      tasks: [],
      filter: 'all'
    }
  },

  watch: {
    activeTabName (val) {
      this.$router.push({ name: val })
    },

    filter () {
      this.fetchTasks()
    }
  },

  created () {
    if (this.$store.state.token === null) {
      this.$router.push({ name: 'signin' })
    }

    this.fetchTasks()
  },

  methods: {
    fetchTasks () {
      this.$store.dispatch('getTasks', this.filter).then(response => {
        this.tasks = response.data
      })
    },

    addTask () {
      this.$store.dispatch('addTask', { task: this.task }).then(() => {
        this.task = null
      }).finally(() => {
        this.fetchTasks()
      })
    },

    doneTask (task) {
      this.$store.dispatch('doneTask', task.id).then(() => {
        this.fetchTasks()
      })
    },

    deleteTask (task) {
      this.$store.dispatch('deleteTask', task.id).then(() => {
        this.fetchTasks()
      })
    }
  }
}
</script>
