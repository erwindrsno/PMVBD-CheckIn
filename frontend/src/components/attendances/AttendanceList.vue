<script setup>
import { ref, onMounted } from 'vue';
import DeleteConfirmDialog from './DeleteConfirmDialog.vue'

const selectedEvent = ref(null);
const events = ref([]); // You need to populate this from your /api/v1/events endpoint
const attendanceList = ref([]);
const itemToDelete = ref(null);
const deleteDialog = ref(false);

const headers = [
  { title: '#', key: 'idx' },
  { title: 'Nama Peserta', key: 'attendee_name' },
  { title: 'Sekolah', key: 'school_name' },
  { title: 'Kelas', key: 'full_grade' },
  { title: 'Direkam pada', key: 'scanned_at' },
  { title: 'Aksi', key: 'action',align: 'center' }
];

// 1. Fetch available events for the dropdown
const fetchEvents = async () => {
  const res = await fetch('http://localhost:8080/api/v1/events',{
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${sessionStorage.getItem('token')}`
    }
  });
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
  const res = await fetch(`http://localhost:8080/api/v1/attendances/${selectedEvent.value}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${sessionStorage.getItem('token')}`
    }
  });
  if (res.status === 401) {
    sessionStorage.removeItem('token');
    window.location.href = '/login'; // Or use router.push('/login')
    return;
  }
  const attendanceData = await res.json();
  attendanceList.value = attendanceData.data.atvis || [];
};


const openDeleteDialog = (item) => {
  itemToDelete.value = item;
  deleteDialog.value = true;
};

const confirmDelete = async () => {
  if (!itemToDelete.value) return;
  try {
    const response = await fetch(`http://localhost:8080/api/v1/attendances/${itemToDelete.value.attendance_id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${sessionStorage.getItem('token')}`
      }
    });
  if (response.status === 401) {
    sessionStorage.removeItem('token');
    window.location.href = '/login'; // Or use router.push('/login')
    return;
  }
    if (response.ok) {
      await fetchAttendance();
    }
  } catch (error) {
    console.error('Error deleting attendance:', error);
  } finally {
    deleteDialog.value = false;
    itemToDelete.value = null;
  }
};

onMounted(() => {
  fetchEvents();
});
</script>

<template>
  <v-container>
    <h2 class="text-h5 mb-4">Riwayat Kehadiran</h2>

    <!-- Event Selector -->
    <v-select
      v-model="selectedEvent"
      :items="events"
      item-title="name"
      item-value="id"
      label="Filter berdasarkan Agenda"
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

      <!-- Action Column -->
      <template v-slot:header.action>
        <div class="text-center">Aksi</div>
      </template>
      <template v-slot:item.action="{ item }">
        <v-btn icon variant="text" size="small" color="error" @click="openDeleteDialog(item)">
          <v-icon>mdi-delete</v-icon>
        </v-btn>
      </template>
    </v-data-table>
  </v-container>

  <DeleteConfirmDialog
    v-model="deleteDialog"
    :attendee-name="itemToDelete?.attendee_name || ''"
    @confirm="confirmDelete"
  />

</template>
