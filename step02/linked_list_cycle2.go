/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func detectCycle(head *ListNode) *ListNode {
	if head == nil {
			return nil
	}
	
	rabbit := head
	turtle := head
	toggle := true
	for {
			if rabbit.Next == nil || rabbit.Next.Next == nil {
					return nil
			}
			if turtle.Next == nil {
					return nil
			}
			
			if toggle {
					rabbit = rabbit.Next.Next                
			} else {
					turtle = turtle.Next   
			}       
			toggle = !toggle    
			
			if rabbit == turtle {
					return turtle
			}   
	} 
}