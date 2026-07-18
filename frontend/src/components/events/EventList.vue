<script setup>
import { ref, onMounted } from 'vue';
import AddEventDialog from './AddEventDialog.vue';
import StatusConfirmDialog from './StatusConfirmDialog.vue';
import DeleteConfirmDialog from './DeleteConfirmDialog.vue';
import { useEventStatus } from '../../composables/useEventStatus.js';
import {
  fetchEvents as apiFetchEvents,
  saveEvent as apiSaveEvent,
  updateEventStatus as apiUpdateEventStatus,
  deleteEvent as apiDeleteEvent
} from '../../services/eventServices.js';

// --- State ---
const dialog = ref(false);
const loading = ref(false);
const events = ref([]);
const newEvent = ref({ name: '' });

const deleteDialog = ref(false);
const itemToDelete = ref(null);

// --- Composable ---
const { statusDialog, itemToChange, statusMessage, statusMap, openStatusDialog } = useEventStatus();

const headers = [
  { title: '#', key: 'idx' }, // Uses index slot
  { title: 'Nama', key: 'name' },
  { title: 'Status', key: 'status' },
  { title: 'Dibuat pada', key: 'created_at' },
  { title: 'Dimulai pada', key: 'started_at' },
  { title: 'Aksi', key: 'action', align: 'center' }
];

// --- API Logic ---
const fetchEvents = async () => {
  loading.value = true;
  try {
    const result = await apiFetchEvents();
    // Use optional chaining and nullish coalescing for safety
    events.value = result?.data?.events ?? (Array.isArray(result) ? result : []);
  } catch (error) {
    console.error('Error fetching events:', error);
    events.value = []; // Ensure it's an array so the table doesn't break
  } finally {
    loading.value = false; // This MUST always run
  }
};

const saveEvent = async () => {
  try {
    const success = await apiSaveEvent(newEvent.value);
    if (success) {
      await fetchEvents(); // Refresh list
      dialog.value = false;
      newEvent.value = { name: '' };
    }
  } catch (error) {
    console.error('Error saving event:', error);
  }
};

// The action after confirming
const confirmStatusChange = async () => {
  if (!itemToChange.value) return;
  try {
    const nextStatus = itemToChange.value.status === 1 ? 2 : 3; // Example promotion transition
    const success = await apiUpdateEventStatus(itemToChange.value.id, nextStatus);
    if (success) {
      await fetchEvents();
    }
  } catch (error) {
    console.error('Error changing status:', error);
  } finally {
    statusDialog.value = false;
    itemToChange.value = null; // Reset
  }
};

const openDeleteDialog = (item) => {
  itemToDelete.value = item;
  deleteDialog.value = true;
};

const confirmDelete = async () => {
  if (!itemToDelete.value) return;
  try {
    const success = await apiDeleteEvent(itemToDelete.value.id);
    if (success) {
      await fetchEvents();
    }
  } catch (error) {
    console.error('Error deleting event:', error);
  } finally {
    deleteDialog.value = false;
    itemToDelete.value = null;
  }
};

const editEvent = (item) => {
  console.log('Edit event:', item);
};

// --- Lifecycle ---
onMounted(() => {
  fetchEvents();
});
</script>

<template>
  <div class="d-flex justify-space-between align-center mb-4">
    <h1>Event List</h1>
    <v-btn color="primary" @click="dialog = true">Tambah Agenda Baru</v-btn>
  </div>

  <v-data-table
    :headers="headers"
    :items="events"
    :loading="loading"
    class="elevation-1"
  >
    <template v-slot:header.status>
      <div class="text-center">Status</div>
    </template>
    <!-- Incrementing Index -->
    <template v-slot:item.idx="{ index }">
      {{ index + 1 }}
    </template>

    <template v-slot:item.status="{ item }">
      <div class="d-flex justify-center">
        <v-chip
          :color="statusMap[item.status]?.color || 'grey'"
          size="small"
          variant="tonal"
        >
          {{ statusMap[item.status]?.label || 'Unknown' }}
        </v-chip>
      </div>
    </template>

    <!-- Refactored Action Column -->
    <template v-slot:item.action="{ item }">
      <div class="d-flex justify-center gap-2">
        <v-btn
          icon
          variant="text"
          size="small"
          color="primary"
          @click="openStatusDialog(item)"
          :disabled="item.status === 3"
        >
          <v-icon>mdi-cog-refresh</v-icon>
        </v-btn>
        <v-btn
          disabled
          icon
          variant="text"
          size="small"
          color="primary"
          @click="editEvent(item)"
        >
          <v-icon>mdi-pencil</v-icon>
        </v-btn>
        <v-btn
          icon
          variant="text"
          size="small"
          color="error"
          @click="openDeleteDialog(item)"
        >
          <v-icon>mdi-delete</v-icon>
        </v-btn>
      </div>
    </template>
  </v-data-table>

  <!-- Add Event Dialog -->
  <AddEventDialog
    v-model="dialog"
    :new-event="newEvent"
    @save="saveEvent"
  />

  <!-- Status Confirm Dialog -->
  <StatusConfirmDialog
    v-model="statusDialog"
    :status-message="statusMessage"
    @confirm="confirmStatusChange"
  />

  <!-- Delete Confirm Dialog -->
  <DeleteConfirmDialog
    v-model="deleteDialog"
    :event-name="itemToDelete?.name || ''"
    @confirm="confirmDelete"
  />
</template>
