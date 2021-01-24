/**
 * Sort function to sort array of object based on selected criteria in ascending or 
 * descending order. Parameters are partially applied. Does not mutate original array.
 * @param criteria => Object attribute used for comparison
 * @param ascending => Boolean value to define asc/desc, true is ascending
 * @param array => Object array to be sorted
 */
const sortBy = criteria => ascending => array => 
    [...array].sort((a,b) => a[criteria] < b[criteria]
        ? (ascending ? 1 : -1) * -1 
        : (ascending ? 1 : -1) * 1)

export default sortBy