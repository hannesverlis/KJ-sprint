package sprint

func Sqrt(n int) int {
    if n < 0 {
        return 0 
    } else if n == 0 || n == 1 {
        return n 
    }
   // Binary search for the square root
   low, high := 1, n
   for low <= high {
	   mid := low + (high-low)/2 // Avoid overflow
	   square := mid * mid

	   if square == n {
		   return mid
	   } else if square < n {
		   low = mid + 1
	   } else {
		   high = mid - 1
	   }
   }

   return 0 
}