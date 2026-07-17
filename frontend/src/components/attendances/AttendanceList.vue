<script setup>
import { ref, onMounted } from 'vue';

const selectedEvent = ref(null);
const events = ref([]); // You need to populate this from your /api/v1/events endpoint
const attendanceList = ref([]);

const headers = [
  { title: 'No.', key: 'idx' },
  { title: 'Attendee Name', key: 'attendee_name' },
  { title: 'School', key: 'school_name' },
  { title: 'Grade', key: 'full_grade' },
  { title: 'Scanned At', key: 'scanned_at' },
];

// 1. Fetch available events for the dropdown
const fetchEvents = async () => {
  const res = await fetch('http://localhost:8080/api/v1/events');
  const eventsListData = await res.json();
  events.value = eventsListData.data.events;
};

// 2. Fetch attendance based on selected event
const fetchAttendance = async () => {
  if (!selectedEvent.value) {
    attendanceList.value = [];
    return;
  }

  // Update this to match the route you fixed (Path or Query)
  const res = await fetch(`http://localhost:8080/api/v1/attendances/${selectedEvent.value}`);
  const attendanceData = await res.json();
  attendanceList.value = attendanceData.data.atvis || [];
};

onMounted(() => {
  fetchEvents();
});
</script>

<template>
  <v-container>
    <h2 class="text-h5 mb-4">Attendance Records</h2>

    <!-- Event Selector -->
    <v-select
      v-model="selectedEvent"
      :items="events"
      item-title="name"
      item-value="id"
      label="Filter by Event"
      variant="outlined"
      class="mb-4"
      clearable
      @update:model-value="fetchAttendance"
    ></v-select>

    <v-data-table
      :headers="headers"
      :items="attendanceList"
      class="elevation-1"
      density="comfortable"
    >
      <template v-slot:item.idx="{ index }">
        {{ index + 1 }}
      </template>
      <template v-slot:item.full_grade="{ item }">
        {{ item.grade_label }} - {{ item.subgrade_name }}
      </template>
      <template v-slot:item.scanned_at="{ item }">
        {{ new Date(item.scanned_at).toLocaleString() }}
      </template>
    </v-data-table>
  </v-container>
</template>
