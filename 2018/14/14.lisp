; 14 of AoC

(defparameter *loops* 47801)

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



