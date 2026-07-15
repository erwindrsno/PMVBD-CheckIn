<script setup>
  import { ref, onMounted } from 'vue';

  // --- State ---
  const dialog = ref(false);
  const detailDialog = ref(false); // Add this
  const loading = ref(false);
  const attendees = ref([]);
  const selectedAttendee = ref(null);
  const newAttendee = ref({ name: '' });

  const headers = [
    { title: 'No.', key: 'idx', width: '50px' },
    { title: 'Name', key: 'name', width: '350px' },
    { title: 'School', key: 'school', width: '300px' },
    { title: 'Grade', key: 'grade', width: '200px' },
    { title: 'Telp', key: 'telp' },
    { title: 'Action', key: 'action', width:'200px' ,align: 'center' }
  ];

  const viewDetails = (item) => {
    selectedAttendee.value = item;
    detailDialog.value = true;
  };

  const fetchAttendees = async () => {
    loading.value = true;
    try {
      // Replace with your actual events endpoint
      const response = await fetch('http://localhost:8080/api/v1/attendees');
      const result = await response.json();
      console.log(result)
      attendees.value = result.data.attendees || [];
    } catch (error) {
      console.error('Error fetching attendees:', error);
    } finally {
      loading.value = false;
    }
  };


  // --- Lifecycle ---
  onMounted(() => {
    fetchAttendees();
  });
</script>

<template>
  <div class="d-flex justify-space-between align-center mb-4">
    <h1>Attendee List</h1>
    <v-btn color="primary" @click="dialog = true">Add New Attendee</v-btn>
  </div>

  <v-data-table
    :headers="headers"
    :items="attendees"
    :loading="loading"
    class="elevation-1"
  >
    <!-- Index -->
    <template v-slot:item.idx="{ index }">
      {{ index + 1 }}
    </template>

    <template v-slot:item.grade="{ item }">
      {{ item.grade }} - {{ item.subgrade }}
    </template>

    <!-- Action Column -->
    <template v-slot:header.action>
      <div class="text-center">Action</div>
    </template>
    <template v-slot:item.action="{ item }">
      <!-- Call viewDetails(item) instead of editEvent -->
      <v-btn icon variant="text" size="small" color="info" @click="viewDetails(item)">
        <v-icon>mdi-eye</v-icon>
      </v-btn>

      <v-btn disabled icon variant="text" size="small" color="primary" @click="editEvent(item)">
        <v-icon>mdi-pencil</v-icon>
      </v-btn>

      <v-btn icon variant="text" size="small" color="error" @click="deleteEvent(item.id)">
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </template>
  </v-data-table>

  <!-- View Detail Dialog -->
  <v-dialog v-model="detailDialog" max-width="600px">
    <v-card v-if="selectedAttendee">
      <v-card-title class="bg-primary text-white">Attendee Details</v-card-title>
      <v-card-text class="mt-4">
        <v-list density="compact">
          <v-list-item title="Name" :subtitle="selectedAttendee.name"></v-list-item>
          <v-list-item title="ID" :subtitle="selectedAttendee.public_id"></v-list-item>
          <v-list-item title="School" :subtitle="selectedAttendee.school"></v-list-item>
          <v-list-item title="Grade" :subtitle="`${selectedAttendee.grade} - ${selectedAttendee.subgrade}`"></v-list-item>
          <v-list-item title="Phone" :subtitle="selectedAttendee.telp"></v-list-item>
          <v-list-item title="Guardian Phone" :subtitle="selectedAttendee.guardian_telp"></v-list-item>
        </v-list>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="primary" @click="detailDialog = false">Close</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
