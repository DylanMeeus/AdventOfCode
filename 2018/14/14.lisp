; 14 of AoC

(defparameter *loops* 47801)
(defparameter *lookup* '(0 4 7 8 0 1))

; get the digits of a number
(defun digits(n) 
  (map 'list #'digit-char-p (prin1-to-string n)))


(defun next-index(lst pos score)
  (mod (+ pos (1+ score)) (list-length lst)))

; solve one1
(defun sum(input recipes fst snd sumlist)
  ; just test with adding one for starters
  (let* ((x (nth fst input))
        (y (nth snd input))
        (sm (digits (+ x y))))
  (setf input (append input sm))
  (setf sumlist (append sumlist sm))
  (setf recipes (+ recipes (list-length sm)))
  (if (>= (list-length sumlist) 10)
    (subseq sumlist 0 10)
    (sum input recipes (next-index input fst x) (next-index input snd y) sumlist))
  ))


; generate recipes
(defun generate(input recipes fst snd)
  ; just test with adding one for starters
  (let* ((x (nth fst input))
        (y (nth snd input))
        (sm (digits (+ x y))))
  (setf input (append input sm))
  (setf recipes (+ recipes (list-length sm)))
  (if (>= recipes *loops*)
    (sum input recipes (next-index input fst x) (next-index input snd y) (subseq (reverse input) 0 (- recipes *loops*)))
    (generate input recipes (next-index input fst x) (next-index input snd y))
  )))

; solve part 2
(defun solve2(input recipes fst snd)
  ; just test with adding one for starters
  (let* ((x (nth fst input))
        (y (nth snd input))
        (sm (digits (+ x y)))
        (index (contains (if (< (list-length input) 20)
                           (reverse input)
                           (subseq (reverse input) 0 10)) (reverse *lookup*) 0))
        )
  (setf input (append input sm))
  (setf recipes (+ recipes (list-length sm)))
  (print recipes)

  ; check for containing substring
  (if (= index -1)
    (solve2 input recipes (next-index input fst x) (next-index input snd y))
    (contains input *lookup* 0)
    )))

; return index of needle in list
(defun contains(lst needle index) 
    (let ((len (list-length needle)))
      (if (or (eq lst nil) (< (list-length lst) len))
        -1
        (progn
          (if (not(eq (search (subseq lst 0 len) needle) nil))
            index
            (contains (cdr lst) needle (1+ index))
            )))))



; solve part 1
; (generate '(3 7) 2 0 1)
; solve part 2
(defun main()
    (solve2 '(3 7) 2 0 1))

