package sprint


func RemoveDuplicates(slice []int) []int {
   // Create a map to store unique elements
   seen := make(map[int]bool)
   result := []int{}
   
   // Loop through the slice, adding elements to the map if they haven't been seen before
   for _, val := range slice {
      if _, ok := seen[val]; !ok {
         seen[val] = true
         result = append(result, val)
      }
   }  
   return result
}
