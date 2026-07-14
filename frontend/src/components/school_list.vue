<script setup>
import { ref } from 'vue';

const dialog = ref(false); // Controls the dialog visibility
const headers = [
  { title: 'No.', key: 'id' },
  { title: 'School Name', key: 'name' },
  { title: 'Location', key: 'location' },
  { title: 'Contact', key: 'contact' }
];
const schools = ref([
  { id: '1', name: 'SMA Negeri 1 Karimun', location: 'Karimun', contact: '0812...' }
]);

// New school form state
const newSchool = ref({ name: '', location: '', contact: '' });

const saveSchool = () => {
  // Add logic to save to your Go backend here later
  schools.value.push({ ...newSchool.value });
  dialog.value = false; // Close dialog
  newSchool.value = { name: '', location: '', contact: '' }; // Reset form
};
</script>

<template>
  <div class="d-flex justify-space-between align-center mb-4">
    <h1>School List</h1>
    <v-btn color="primary" @click="dialog = true">Add New School</v-btn>
  </div>

  <v-data-table :headers="headers" :items="schools" :items-per-page="5" class="elevation-1"></v-data-table>

  <!-- Add School Dialog -->
  <v-dialog v-model="dialog" max-width="500px">
    <v-card>
      <v-card-title>Add New School</v-card-title>
      <v-card-text>
        <v-text-field v-model="newSchool.name" label="School Name"></v-text-field>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="grey" @click="dialog = false">Cancel</v-btn>
        <v-btn color="primary" @click="saveSchool">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
