; first puzzle of AoC 2018

(defun get-file()
  (with-open-file (stream "input.txt")
    (loop for line = (read-line stream nil)
          while line
          collect line)))

(defun get-test()
  '("1" "-2" "3" "1"))

(defun solve() 
  (let ((input (mapcar (lambda (s) (parse-integer s)) (get-file))))
    (apply '+ input)))


(defun solve2() 
  (let ((input (mapcar (lambda (s) (parse-integer s)) (get-file))))
    (next (car input) (cdr input) '(0))))

(defun next(f fs pfs)
  (let ((next_value (+ f (car fs))))
    (cond ((eq next_value (find next_value pfs))
           (write next_value))
          ((not (eq nil (cdr fs)))
            (next next_value (cdr fs) (cons next_value pfs)))
          (t (next next_value (mapcar (lambda (s) (parse-integer s)) (get-file)) (cons next_value pfs))))))
